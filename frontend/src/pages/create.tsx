import React from 'react';
import Head from 'next/head';
import { Flex, useToast } from '@chakra-ui/react';
import GIFCreation from '@/components/Forms/GIFCreation';
import { useMutation } from '@apollo/client';
import { useRouter } from 'next/router';

import { createGIFMutation } from '@/services/mutations';

export default function Search() {
  const toast = useToast();
  const router = useRouter();

  const [createGIF] = useMutation(createGIFMutation, {
    onCompleted: () => {
      toast({
        title: 'GIF creado exitosamente',
        status: 'success',
        isClosable: true,
      });

      router.push(router.basePath);
    },
    onError: ({ message }) => {
      toast({
        title: 'Error en la creación del GIF',
        description: `Detalle: ${message}`,
        status: 'error',
        isClosable: true,
      });
    },
  });
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
          <GIFCreation
            submitFile={(newGIF) => createGIF({
              variables: {
                input: {
                  name: newGIF.name,
                  content: newGIF.content,
                  user: newGIF.user,
                  tags: newGIF.tags,
                },
              },
            })}
          />
        </Flex>
      </main>
    </>
  );
}
