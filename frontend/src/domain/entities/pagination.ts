export type Pagination<T> = {
  page: number
  rowsPerPage: number
  totalRows: number
  results: T[]
};

export const makePagination = <T>(
  input: Partial<Pagination<T>>,
): Pagination<T> => {
  let page: number = input.page ?? 0;
  let rowsPerPage: number = input.rowsPerPage ?? 0;
  let totalRows: number = input.totalRows ?? 0;
  let results: T[] = input.results ?? [];

  return {
    get page(): number {
      return page;
    },
    set page(newPage: number) {
      page = newPage;
    },
    get rowsPerPage(): number {
      return rowsPerPage;
    },
    set rowsPerPage(newRowsPerPage: number) {
      rowsPerPage = newRowsPerPage;
    },
    get totalRows(): number {
      return totalRows;
    },
    set totalRows(newTotalRows: number) {
      totalRows = newTotalRows;
    },
    get results(): T[] {
      return results;
    },
    set results(newResults: T[]) {
      results = newResults;
    },
  };
};
