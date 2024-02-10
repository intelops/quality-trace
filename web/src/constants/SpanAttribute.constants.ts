import {SemanticAttributes, SemanticResourceAttributes} from '@opentelemetry/semantic-conventions';

export const TraceTestAttributes = {
  NAME: 'name',
  KIND: 'quality-trace.span.kind',
  QUALITYTRACE_SPAN_TYPE: 'quality-trace.span.type',
  QUALITYTRACE_SPAN_DURATION: 'quality-trace.span.duration',
  QUALITYTRACE_RESPONSE_STATUS: 'quality-trace.response.status',
  QUALITYTRACE_RESPONSE_BODY: 'quality-trace.response.body',
  QUALITYTRACE_RESPONSE_HEADERS: 'quality-trace.response.headers',
  QUALITYTRACE_SELECTED_SPANS_COUNT: 'quality-trace.selected_spans.count',
};

export const Attributes: Record<string, string> = {
  ...SemanticAttributes,
  ...SemanticResourceAttributes,
  ...TraceTestAttributes,
  HTTP_REQUEST_HEADER: 'http.request.header.',
  HTTP_RESPONSE_HEADER: 'http.response.header',
};

export * from '@opentelemetry/semantic-conventions';
