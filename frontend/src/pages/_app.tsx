import React from 'react';
import { ChakraProvider } from '@chakra-ui/react';
import Head from 'next/head';
import { ApolloProvider } from '@apollo/client';
import type { AppProps } from 'next/app';

import { client } from '@/services/graphqlClient';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <meta name="description" content="Página del escritor Andrei Taiba" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <ApolloProvider client={client}>
          <ChakraProvider>
            <Component {...pageProps} />
          </ChakraProvider>
        </ApolloProvider>
      </main>
    </>
  );
}
