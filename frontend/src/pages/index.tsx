import React, { useEffect } from 'react';
import Head from 'next/head';
import { useLazyQuery } from '@apollo/client';
import { CircularProgress, Flex, useToast } from '@chakra-ui/react';

import GIFCard from '@/components/GIFCard';
import { getGIFsWithFilters } from '@/services/queries';
import { Pagination } from '@/domain/entities/pagination';
import { GIF } from '@/domain/entities/gif';

const example: GIF = {
  tags: ['testing', 'hello', 'world'],
  URL: 'https://i.pinimg.com/originals/11/0e/7c/110e7c1e1c8c8953e787b56fdff866ed.gif',
  name: 'Test',
  content: '',
  user: '',
  createdAt: new Date(),
};

export default function Home() {
  const toast = useToast();
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
            page: 1,
            query: '',
          },
        },
      },
    );
  }, [getGIFs]);

  if (loading || error !== undefined) {
    return (
      <>
        <Head>
          <title>Giphy Challenge</title>
        </Head>
        <main>
          <Flex
            width="100wh"
            height="100vh"
            justifyContent="center"
            alignItems="center"
          >
            <CircularProgress isIndeterminate color="green.300" />
          </Flex>
        </main>
      </>
    );
  }

  const { gifs: { results } } = data || { gifs: { results: [example] } };

  return (
    <>
      <Head>
        <title>Giphy Challenge</title>
      </Head>
      <main>
        {results.map((gif) => (
          <GIFCard
            key={gif.name}
            url={gif.URL}
            tags={gif.tags}
            name={gif.name}
          />
        ))}
      </main>
    </>
  );
}
