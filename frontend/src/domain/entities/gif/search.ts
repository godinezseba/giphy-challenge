export type Search = {
  query: string
  page: number
};

export const makeSearch = (
  input: Search,
): Search => {
  let page: number = input.page ?? 0;
  let query: string = input.query ?? '';

  return {
    get page(): number {
      return page;
    },
    set page(newPage: number) {
      page = newPage;
    },
    get query(): string {
      return query;
    },
    set query(newquery: string) {
      query = newquery;
    },
  };
};
