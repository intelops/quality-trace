import {render} from 'test-utils';
import TestMock from '../../../models/__mocks__/Test.mock';
import TestHeader from '../TestHeader';

test('SpanAttributesTable', () => {
  const test = TestMock.model();

  const onDelete = jest.fn;
  const onEdit = jest.fn;
  const shouldEdit = true;

  const {getByTestId} = render(
    <TestHeader description={test.description} id={test.id} shouldEdit={shouldEdit} onEdit={onEdit} onDelete={onDelete} title={test.name} runButton={<div />} />
  );
  expect(getByTestId('test-details-name')).toBeInTheDocument();
});
