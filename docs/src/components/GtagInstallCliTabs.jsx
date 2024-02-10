import React from 'react';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from '@theme/CodeBlock';

export default function GtagInstallCliTabs() {
  function trigger({ event, installationPlatform }) {
    window.dataLayer = window.dataLayer || [];
    window.dataLayer.push({
    'installationPlatform': installationPlatform,
    'event': 'installationPlatform',
    });
  }
  return (
    <Tabs groupId="operating-systems">
      <TabItem value="mac" label="MAC" default>
        <div onClick={() => trigger({ installationPlatform: 'MacOS' })}>
          <CodeBlock
              language="bash"
              title="Terminal"
              >
            {`brew install intelops/quality-trace/quality-trace`}
          </CodeBlock>
        </div>
      </TabItem>
      <TabItem value="linux" label="LINUX">
        <div onClick={() => trigger({ installationPlatform: 'Linux' })}>
          <CodeBlock
              language="bash"
              title="Terminal"
          >
          {`curl -L https://raw.githubusercontent.com/intelops/quality-trace/main/install-cli.sh | bash`}
          </CodeBlock>
        </div>
      </TabItem>
      <TabItem value="win" label="WINDOWS">
        <div onClick={() => trigger({ installationPlatform: 'Windows' })}>
          <CodeBlock
              language="bash"
              title="Terminal"
          >
          {`choco source add --name=intelops_repo --source=https://chocolatey.intelops.io/chocolatey ; choco install quality-trace`}
          </CodeBlock>
        </div>
      </TabItem>
    </Tabs>
  );
};
