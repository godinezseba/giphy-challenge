import React from 'react';
import {
  Badge,
} from '@chakra-ui/react';

type TagProps = {
  name: string
};

export default function Tag(props: TagProps) {
  const {
    name,
  } = props;

  return (
    <Badge rounded="full" px="2" fontSize="0.8em" colorScheme="red">
      {name}
    </Badge>
  );
}
