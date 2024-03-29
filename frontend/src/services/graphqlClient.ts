import { ApolloClient, InMemoryCache } from '@apollo/client';

export const client = new ApolloClient({
  uri: process.env.NEXT_PUBLIC_BFF_URL,
  cache: new InMemoryCache(),
});

export default client;
