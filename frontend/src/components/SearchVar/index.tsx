import React, { useState } from 'react';
import { SearchIcon } from '@chakra-ui/icons';
import { Button, Flex, Input } from '@chakra-ui/react';

type SearchVarProps = {
  actualQuery: string
  setQuery: (newQuery: string) => void
  isLoading: boolean
};

export default function SearchVar(props: SearchVarProps) {
  const {
    actualQuery,
    isLoading,
    setQuery,
  } = props;
  const [inputQuery, setInputQuery] = useState(actualQuery);

  return (
    <Flex>
      <Input
        placeholder="Busca tus GIFs favoritos!"
        value={inputQuery}
        onChange={(event) => setInputQuery(event.target.value)}
        maxW={500}
      />
      <Button
        rightIcon={<SearchIcon />}
        isLoading={isLoading}
        onClick={() => { setQuery(inputQuery); }}
      >
        Buscar
      </Button>
    </Flex>
  );
}
