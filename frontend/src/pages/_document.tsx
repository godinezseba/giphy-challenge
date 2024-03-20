import React from 'react';
import {
  Html,
  Head,
  Main,
  NextScript,
} from 'next/document';

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link rel="stylesheet" type="text/css" href="//fonts.googleapis.com/css?family=EB+Garamond" />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
