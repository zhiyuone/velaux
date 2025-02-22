import React, { Component } from 'react';
import { If } from 'tsx-control-statements/components';
import { connect } from 'dva';
import WorkflowComponent from './workflow-component';
import type { WorkFlowData } from './entity';
import { getWorkFlowDefinitions } from '../../api/workflows';
import type { ApplicationDetail } from '../../interface/application';
import './index.less';

type Props = {
  workflowList: WorkFlowData[];
  applicationDetail?: ApplicationDetail;
  dispatch: ({}) => {};
  match: {
    params: {
      appName: string;
    };
  };
  history: {
    push: (path: string, state: {}) => {};
  };
};

type State = {
  appName: string;
  workFlowDefinitions: [];
};
@connect((store: any) => {
  return { ...store.workflow, ...store.application };
})
class Workflow extends Component<Props, State> {
  constructor(props: any) {
    super(props);
    const { params } = this.props.match;
    this.state = {
      appName: params.appName || '',
      workFlowDefinitions: [],
    };
  }

  componentDidMount() {
    this.onGetWorkflow();
    this.onGetWorkflowDefinitions();
  }

  onGetWorkflow = () => {
    this.props.dispatch({
      type: 'workflow/getWorkflowList',
      payload: {
        appName: this.state.appName,
      },
    });
  };

  onGetWorkflowDefinitions = async () => {
    getWorkFlowDefinitions().then((res: any) => {
      if (res) {
        this.setState({
          workFlowDefinitions: res && res.definitions,
        });
      }
    });
  };

  render() {
    const { workflowList, dispatch, applicationDetail } = this.props;
    const { params } = this.props.match;
    return (
      <div style={{ height: '100%' }} className="workflow-wraper">
        <If condition={workflowList.length > 0}>
          <React.Fragment>
            {workflowList.map((workflow: WorkFlowData) => (
              <WorkflowComponent
                key={workflow.name + params.appName + workflow.steps}
                appName={params.appName}
                data={workflow}
                workFlowDefinitions={this.state.workFlowDefinitions}
                getWorkflow={this.onGetWorkflow}
                applicationDetail={applicationDetail}
                dispatch={dispatch}
              />
            ))}
          </React.Fragment>
        </If>
      </div>
    );
  }
}

export default Workflow;
