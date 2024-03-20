import React from 'react';
import {
  Flex,
  Box,
  Image,
} from '@chakra-ui/react';
import Tag from '@/components/Tag';

type GIFCardProps = {
  url: string
  tags: string[]
  name: string
};

export default function GIFCard(props: GIFCardProps) {
  const {
    url,
    tags,
    name,
  } = props;

  return (
    <Flex p={50} w="full" alignItems="center" justifyContent="center">
      <Box
        bg="white"
        maxW="sm"
        borderWidth="1px"
        rounded="lg"
        shadow="lg"
        position="relative"
      >
        <Image src={url} alt={name} roundedTop="lg" />

        <Box p="6">
          <Box display="flex" alignItems="baseline">
            {tags.map((tag) => (
              <Tag key={`${name}-${tag}`} name={tag} />
            ))}
          </Box>
        </Box>
      </Box>
    </Flex>
  );
}
