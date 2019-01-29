'use strict';

const _ = require('lodash');
const MeterInstanceJob = require('../../jobs/MeterInstanceJob');
const CONST = require('../../common/constants');

const meterGuid = 'meter-guid';
const not_excluded_plan = 'bc158c9a-7934-401e-94ab-057082a5073f';
let options_json = {
  id: meterGuid,
  timestamp: '2019-01-21T11:00:42.384518Z',
  service: {
    id: '24731fb8-7b84-4f57-914f-c3d55d793dd4',
    plan: not_excluded_plan
  },
  consumer: {
    environment: '',
    region: '',
    org: '33915d88-6002-4e83-b154-9ec2075e1435',
    space: 'bd78dbbb-5225-4dfa-94e0-816a4de9b7c9',
    instance: '4e099918-1b37-42a8-9dbd-a752225fcd07'
  },
  measues: [{
    id: 'instances',
    value: 'start'
  }]
};

function getDummyEvent(options_json) {
  let dummy_event = {
    apiVersion: 'instance.servicefabrik.io/v1alpha1',
    kind: 'Sfevent',
    metadata: {
      clusterName: '',
      creationTimestamp: '2019-01-21T11:00:43Z',
      generation: 1,
      labels: {
        state: 'TO_BE_METERED'
      },
      name: meterGuid,
      namespace: 'default',
      resourceVersion: '326999',
      selfLink: '/apis/instance.servicefabrik.io/v1alpha1/namespaces/default/sfevents/48eb2e1e-dbfa-4554-9663-273418437e90',
      uid: 'cbf0c777-1d6b-11e9-806d-0e655bfa3b31'
    },
    spec: {
      metadata: {
        creationTimestamp: null
      },
      options: options_json
    }
  };
  return dummy_event;
}

describe('Jobs', () => {
  describe('MeterInstanceJobs', () => {
    describe('#sendEvent', () => {
      it('should not send metering event for excluded plans', () => {
        const expectedResponse = {
          status: 200
        };
        const payload = {
          status: {
            state: CONST.METER_STATE.EXCLUDED
          }
        };
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        // updated the dummy event with exluded plans
        options_json.service.id = '24731fb8-7b84-4f57-914f-c3d55d793dd4';
        options_json.service.plan = '466c5078-df6e-427d-8fb2-c76af50c0f56';
        const dummy_event = getDummyEvent(options_json);
        return MeterInstanceJob.sendEvent(dummy_event)
          .then(res => {
            expect(res).to.eql(true);
            mocks.verify();
          });
      });
      it('should send document to metering and update state in apiserver', () => {
        const expectedResponse = {
          status: 200
        };
        const payload = {
          status: {
            state: CONST.METER_STATE.METERED
          }
        };
        const mock_token = 'mock_token_string';
        const mock_response_code = 200;
        mocks.metering.mockAuthCall(mock_token);
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code, () => {
          return true;
        });
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        // updated the dummy event with not exluded plans
        options_json.service.id = '24731fb8-7b84-4f57-914f-c3d55d793dd4';
        options_json.service.plan = not_excluded_plan;
        const dummy_event = getDummyEvent(options_json);
        return MeterInstanceJob.sendEvent(dummy_event)
          .then(res => {
            expect(res).to.eql(true);
            mocks.verify();
          })
          .catch(err => expect(err).to.be.undefined);
      });
      it('should update state in apiserver if sending document fails', () => {
        const expectedResponse = {
          status: 200
        };
        const payload = {
          status: {
            state: CONST.METER_STATE.FAILED
          }
        };
        const mock_token = 'mock_token_string';
        const mock_response_code = 400;
        mocks.metering.mockAuthCall(mock_token);
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code, () => {
          return true;
        });
        // updated the dummy event with not exluded plans
        options_json.service.id = '24731fb8-7b84-4f57-914f-c3d55d793dd4';
        options_json.service.plan = not_excluded_plan;
        const dummy_event = getDummyEvent(options_json);
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        return MeterInstanceJob.sendEvent(dummy_event)
          .then(res => {
            expect(res).to.eql(false);
            mocks.verify();
          })
          .catch(err => expect(err).to.be.undefined);
      });
    });
    describe('#meter', () => {
      it('Should send event for all documents', () => {
        //Send doucments to metering service 2 times
        const mock_token = 'mock_token_string';
        const mock_response_code = 200;
        const expectedResponse = {
          status: 200
        };
        const payload = {
          status: {
            state: CONST.METER_STATE.METERED
          }
        };
        mocks.metering.mockAuthCall(mock_token);
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code, () => {
          return true;
        });
        mocks.metering.mockAuthCall(mock_token);
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code, () => {
          return true;
        });
        // updated the dummy event with not exluded plans
        options_json.service.id = '24731fb8-7b84-4f57-914f-c3d55d793dd4';
        options_json.service.plan = not_excluded_plan;
        const dummy_event = getDummyEvent(options_json);
        // update apiserver for the 2 events
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        return MeterInstanceJob.meter([dummy_event, _.cloneDeep(dummy_event)])
          .then(res => {
            expect(res.totalEvents).to.eql(2);
            expect(res.success).to.eql(2);
            expect(res.failed).to.eql(0);
          })
          .catch(err => expect(err).to.be.undefined);
      });
      it('Should keep tab of failed events', () => {
        //Send doucments to metering service 2 times
        const mock_token = 'mock_token_string';
        const mock_response_code = 200;
        const mock_response_code_failure = 400;
        const expectedResponse = {
          status: 200
        };
        const payload = {
          status: {
            state: CONST.METER_STATE.METERED
          }
        };
        const payload_failure = {
          status: {
            state: CONST.METER_STATE.FAILED
          }
        };
        mocks.metering.mockAuthCall(mock_token);
        // mock successfull put req
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code, () => {
          return true;
        });
        mocks.metering.mockAuthCall(mock_token);
        // mock failed put req
        mocks.metering.mockSendUsageRecord(mock_token, mock_response_code_failure, () => {
          return true;
        });
        // updated the dummy event with not exluded plans
        options_json.service.id = '24731fb8-7b84-4f57-914f-c3d55d793dd4';
        options_json.service.plan = not_excluded_plan;
        const dummy_event = getDummyEvent(options_json);
        // update apiserver for the 2 events
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload);
        mocks.apiServerEventMesh.nockPatchResource(CONST.APISERVER.RESOURCE_GROUPS.INSTANCE,
          CONST.APISERVER.RESOURCE_TYPES.SFEVENT, meterGuid, expectedResponse, 1, payload_failure);
        return MeterInstanceJob.meter([dummy_event, _.cloneDeep(dummy_event)])
          .then(res => {
            expect(res.totalEvents).to.eql(2);
            expect(res.success).to.eql(1);
            expect(res.failed).to.eql(1);
          });
      });
    });
  });
});