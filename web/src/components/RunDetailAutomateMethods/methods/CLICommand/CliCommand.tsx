import {ReadOutlined} from '@ant-design/icons';
import {FramedCodeBlock} from 'components/CodeBlock';
import {ResourceType} from 'types/Resource.type';
import * as S from './CliCommand.styled';
import Controls from './Controls';
import useCliCommand from './hooks/useCliCommand';

interface IProps {
  id: string;
  variableSetId?: string;
  fileName?: string;
  resourceType: ResourceType;
  docsUrl?: string;
}

const CLiCommand = ({id, variableSetId, fileName = '', resourceType, docsUrl}: IProps) => {
  const {command, onGetCommand} = useCliCommand();

  return (
    <S.Container>
      <S.TitleContainer>
        <S.Title>CLI Configuration</S.Title>
        <a href={docsUrl} target="_blank">
          <ReadOutlined />
        </a>
      </S.TitleContainer>
      <FramedCodeBlock
        title="Qualitytrace CLI command:"
        language="bash"
        value={command}
        minHeight="100px"
        maxHeight="100px"
      />
      <Controls
        onChange={onGetCommand}
        id={id}
        fileName={fileName}
        variableSetId={variableSetId}
        resourceType={resourceType}
      />
    </S.Container>
  );
};

export default CLiCommand;
