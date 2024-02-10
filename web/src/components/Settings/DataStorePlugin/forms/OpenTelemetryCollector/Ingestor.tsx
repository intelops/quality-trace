import {Form, Switch} from 'antd';
import DocsBanner from 'components/DocsBanner/DocsBanner';
import {INGESTOR_ENDPOINT_URL} from 'constants/Common.constants';
import {TCollectorDataStores, TDraftDataStore} from 'types/DataStore.types';
import * as S from './OpenTelemetryCollector.styled';

const Ingestor = () => {
  const form = Form.useFormInstance<TDraftDataStore>();
  const dataStoreType = Form.useWatch('dataStoreType', form) as TCollectorDataStores;
  const baseName = ['dataStore', dataStoreType];

  return (
    <S.Container>
      <S.Description>
        Qualitytrace easily integrates with any distributed tracing solution via the OpenTelemetry Collector. It allows
        your current tracing system to send only Qualitytrace spans while the rest go to your chosen backend.
      </S.Description>
      <S.Title>Ingestor Endpoint</S.Title>
      <S.Description>
        Qualitytrace exposes trace ingestion endpoints on ports 4317 for gRPC and 4318 for HTTP. Turn on the Qualitytrace
        ingestion endpoint to start sending traces. Use the Qualitytrace Server’s hostname and port to connect.For example,
        with Docker use quality-trace:4317 for gRPC.
      </S.Description>
      <S.SwitchContainer>
        <Form.Item name={[...baseName, 'isIngestorEnabled']} valuePropName="checked">
          <Switch />
        </Form.Item>
        Enable Qualitytrace ingestion endpoint
      </S.SwitchContainer>
      <DocsBanner>
        Need more information about setting up ingestion endpoint?{' '}
        <a target="_blank" href={INGESTOR_ENDPOINT_URL}>
          Go to our docs
        </a>
      </DocsBanner>
    </S.Container>
  );
};

export default Ingestor;
