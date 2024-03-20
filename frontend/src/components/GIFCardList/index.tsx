import React from 'react';
import {
  CircularProgress,
  Flex,
} from '@chakra-ui/react';

import GIFCard from '@/components/GIFCard';
import Pagination from '../Pagination';
import SearchVar from '../SearchVar';

type GIFCardListProps = {
  gifs: {
    URL: string
    tags: string[]
    name: string
  }[]
  actualPage: number
  setActualPage: (newPage: number) => void
  totalPages: number
  isLoading: boolean
  actualQuery: string
  setQuery: (newQuery: string) => void
};

export default function GIFCardList(props: GIFCardListProps) {
  const {
    gifs,
    isLoading,
    actualPage,
    totalPages,
    setActualPage,
    actualQuery,
    setQuery,
  } = props;

  return (
    <Flex
      w="full"
      justifyContent="center"
      flexDirection="column"
      paddingY={200}
      paddingX={200}
    >
      <SearchVar
        actualQuery={actualQuery}
        setQuery={setQuery}
        isLoading={isLoading}
      />

      <Pagination
        actualPage={actualPage}
        isLoading={isLoading}
        setActualPage={setActualPage}
        maxPages={totalPages}
      />

      <Flex
        width="full"
        height="full"
        justifyContent="center"
        alignItems="center"
        flexWrap="wrap"
      >
        { isLoading ? (
          <CircularProgress isIndeterminate color="green.300" />
        ) : gifs.map((gif) => (
          <GIFCard
            key={gif.name}
            url={gif.URL}
            tags={gif.tags}
            name={gif.name}
          />
        ))}
      </Flex>

      <Pagination
        actualPage={actualPage}
        isLoading={isLoading}
        setActualPage={setActualPage}
        maxPages={totalPages}
      />
    </Flex>
  );
}
