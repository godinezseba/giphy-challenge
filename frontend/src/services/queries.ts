import { gql } from '@apollo/client';

export const getGIFsWithFilters = gql`
  query GetGIFsWithFilters($searchQuery: GIFSearchInput!){
    gifs(input: $searchQuery) {
      page
      rowsPerPage
      totalRows
      results {
        URL
        tags
        name
      }
    }
  }
`;

export default {
  getGIFsWithFilters,
};
