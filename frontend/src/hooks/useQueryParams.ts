import { useRouter } from 'next/router';
import { useState, useEffect } from 'react';

import { QueryParams } from '@/domain/entities/queryParams';

export const useQueryParams = (): [QueryParams, (newParams: QueryParams) => void] => {
  const router = useRouter();
  const [queryParams, setQueryParams] = useState<QueryParams>({});

  const updateQueryParams = (newParams: QueryParams) => {
    const mergedParams = { ...queryParams, ...newParams };

    setQueryParams(mergedParams);
    router.push({
      pathname: router.pathname,
      query: mergedParams,
    });
  };

  useEffect(() => {
    setQueryParams(router.query);
  }, [router.query]);

  return [queryParams, updateQueryParams];
};

export default useQueryParams;
