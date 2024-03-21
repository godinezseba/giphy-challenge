import React, { useEffect } from 'react';
import Head from 'next/head';
import { useLazyQuery } from '@apollo/client';
import { Flex, useToast } from '@chakra-ui/react';

import { getGIFsWithFilters } from '@/services/queries';
import { Pagination } from '@/domain/entities/pagination';
import { GIF } from '@/domain/entities/gif';
import GIFCardList from '@/components/GIFCardList';
import { useQueryParams } from '@/hooks/useQueryParams';

export default function Home() {
  const toast = useToast();

  const [queryParams, setQueryParams] = useQueryParams();

  const [
    getGIFs,
    { loading, error, data },
  ] = useLazyQuery<{ gifs: Pagination<GIF> }>(
    getGIFsWithFilters,
    {
      onError: ({ message }) => {
        toast({
          title: 'Error en la obtención de las imagenes',
          description: `Detalle: ${message}`,
          status: 'error',
          isClosable: true,
        });
      },
    },
  );

  useEffect(() => {
    getGIFs(
      {
        variables: {
          searchQuery: {
            page: queryParams.page,
            query: queryParams.query,
            pageSize: 9,
          },
        },
      },
    );
  }, [getGIFs, queryParams]);

  const {
    gifs: {
      results = [],
      page = 1,
      totalRows = 1,
      rowsPerPage = 1,
    },
  } = data || { gifs: {} };

  return (
    <>
      <Head>
        <title>Giphy Challenge</title>
      </Head>
      <main>
        <Flex
          w="full"
          flexWrap="wrap"
          justifyContent="center"
          paddingX={10}
        >
          <GIFCardList
            isLoading={!data || loading || error !== undefined}
            gifs={results}
            actualPage={page}
            totalPages={Math.round(totalRows / rowsPerPage) || 1}
            setActualPage={(newPage: number) => {
              setQueryParams({
                page: newPage,
              });
            }}
            actualQuery={queryParams.query || ''}
            setQuery={(newQuery: string) => {
              setQueryParams({
                query: newQuery,
              });
            }}
          />
        </Flex>
      </main>
    </>
  );
}
