import { gql } from '@apollo/client';

export const createGIFMutation = gql`
  mutation CreateGIF($input: GIFInput) {
    createGIF(input: $input) {
      name
      URL
    }
  }
`;

export default createGIFMutation;
