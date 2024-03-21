import { ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons';
import { IconButton, Flex, Text } from '@chakra-ui/react';
import React from 'react';

interface PaginationProps {
  actualPage: number
  maxPages: number
  setActualPage: (newPage: number) => void
  isLoading: boolean
}

export default function Pagination(props: PaginationProps) {
  const {
    actualPage,
    maxPages,
    isLoading,
    setActualPage,
  } = props;

  return (
    <Flex
      w="full"
      justifyContent="space-between"
      alignItems="center"
    >
      <IconButton
        isLoading={isLoading}
        aria-label="Página anterior"
        icon={<ArrowBackIcon />}
        isDisabled={actualPage <= 1}
        onClick={() => setActualPage(actualPage - 1)}
      />

      <Text>
        {`Página ${actualPage} de ${maxPages}`}
      </Text>

      <IconButton
        isLoading={isLoading}
        aria-label="Página siguiente"
        icon={<ArrowForwardIcon />}
        isDisabled={actualPage >= maxPages}
        onClick={() => setActualPage(actualPage + 1)}
      />
    </Flex>
  );
}
