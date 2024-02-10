/* eslint-disable global-require */

import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import Translate from '@docusaurus/Translate';
import Heading from '@theme/Heading';

const GettingStartedGuides = [
  {
    name: '👇 Install Qualitytrace',
    url: './installation',
    description: (
      <Translate >
        Set up Qualitytrace and start trace-based testing your distributed system.
      </Translate>
    ),
    button: 'Set up Qualitytrace',
  },
  {
    name: '🙌 Open Qualitytrace',
    url: './open',
    description: (
      <Translate>
        After installing it, open Qualitytrace start to creating trace-based tests.
      </Translate>
    ),
    button: 'Create tests',
  },
  {
    name: '🤔 Don\'t have OpenTelemetry?',
    url: './no-otel',
    description: (
      <Translate >
        Install OpenTelemetry in 5 minutes without any code changes!
      </Translate>
    ),
    button: 'Set up OTel',
  },
  {
    name: '🤩 Open Source',
    url: 'https://github.com/intelops/quality-trace',
    description: (
      <Translate>
        Check out the Qualitytrace GitHub repo! Please consider giving us a star! ⭐️
      </Translate>
    ),
    button: 'Go to GitHub',
  },
];

interface Props {
  name: string;
  url: string;
  button: string;
  description: JSX.Element;
}

function GettingStartedGuideCard({name, url, description, button}: Props) {
  return (
    <div className="col col--6 margin-bottom--lg">
      <div className={clsx('card')}>
        <div className="card__body">
          <Heading as="h3">{name}</Heading>
          <p>{description}</p>
        </div>
        <div className="card__footer">
          <div className="button-group button-group--block">
            <Link className="button button--secondary" to={url}>
              {button}
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

export function GettingStartedGuideCardsRow(): JSX.Element {
  return (
    <div className="row">
      {GettingStartedGuides.map((gettingStartedGuide) => (
        <GettingStartedGuideCard key={gettingStartedGuide.name} {...gettingStartedGuide} />
      ))}
    </div>
  );
}
