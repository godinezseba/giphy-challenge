import React from 'react';
import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
  useToast,
} from '@chakra-ui/react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';

import { GIF, makeGIF } from '@/domain/entities/gif';
import { toBase64 } from '@/utils/fileToBase64';

const MAX_FILE_SIZE = 10000000;

type FormTypes = {
  name: string
  file: File
  tags: string[]
};

type GIFCreationProps = {
  submitFile: (newGIF: GIF) => Promise<any>
};

const CreationSchema = Yup.object().shape({
  file: Yup.mixed()
    .test({
      message: 'El archivo tiene que ser de tipo GIF',
      test: (file) => (file?.name.includes('gif')),
    })
    .test({
      message: 'El archivo es muy pesado',
      test: (file) => (file?.size < MAX_FILE_SIZE),
    })
    .required('El archivo es requerido'),
  name: Yup.string()
    .min(5, 'Nombre muy corto')
    .max(25, 'Nombre muy largo')
    .required('El nombre es requerido'),
});

export default function GIFCreation(props: GIFCreationProps) {
  const { submitFile } = props;
  const toast = useToast();

  return (
    <Formik
      initialValues={{ name: '', tags: [], file: null!! } as FormTypes}
      onSubmit={async (values, { setSubmitting }) => {
        let file: string = '';
        try {
          file = await toBase64(values.file);
        } catch (error) {
          toast({
            title: 'Error al leer el archivo',
            status: 'error',
            isClosable: true,
          });
          setSubmitting(false);
          return;
        }

        const gif = makeGIF({
          name: values.name,
          content: file,
          tags: values.tags,
        });

        try {
          await submitFile(gif);
        } finally {
          setSubmitting(false);
        }
      }}
      validationSchema={CreationSchema}
    >
      {({
        setFieldValue,
        isSubmitting,
        errors,
        touched,
        isValid,
      }) => (
        <Form>
          <Flex
            direction="column"
            margin={200}
            alignItems="space-around"
          >
            <FormControl isInvalid={!!errors.name && touched.name}>
              <Field
                as={Input}
                id="name"
                name="name"
                placeholder="Nombra tu nuevo GIF"
              />
              <FormErrorMessage>{errors.name}</FormErrorMessage>
            </FormControl>
            <FormControl isInvalid={!!errors.file}>
              <input
                id="file"
                name="file"
                type="file"
                onChange={({ currentTarget }) => {
                  if (!!currentTarget && !!currentTarget.files) {
                    setFieldValue('file', currentTarget.files[0]);
                  }
                }}
              />
              <FormErrorMessage>{errors.file as string}</FormErrorMessage>
            </FormControl>
            <Button
              type="submit"
              isLoading={isSubmitting}
              isDisabled={!isValid}
            >
              Subir
            </Button>
          </Flex>
        </Form>
      )}
    </Formik>
  );
}
