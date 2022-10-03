import { Grid, GridItem, Heading, Link, Stack, Text } from '@chakra-ui/react';
import { FC } from 'react';
import NextLink from 'next/link';

import { CONTRIBUTING_PAGE, DOCS_PAGE, FAQ_PAGE } from '../../../constants';

export const QuickLinks: FC = () => {
  return (
    <Stack border='2px solid #11866f'>
      <Stack p={4} borderBottom='2px solid #11866f'>
        <Heading
          // TODO: move text style to theme
          as='h2'
          fontFamily='"JetBrains Mono", monospace'
          fontWeight={400}
          fontSize='1.5rem'
          lineHeight='auto'
          letterSpacing='4%'
          // TODO: move to theme colors
          color='#1d242c'
        >
          Quick Links
        </Heading>
      </Stack>

      <Grid templateColumns='repeat(2, 1fr)' sx={{ mt: '0 !important' }}>
        {/* get started */}
        <GridItem borderRight='2px solid #11866f' borderBottom='2px solid #11866f'>
          <Stack p={4} h='100%'>
            <Text fontFamily='"Inter", sans-serif' lineHeight='26px' mb={-1}>
              Don&apos;t know where to start?
            </Text>

            <Text fontFamily='"Inter", sans-serif' lineHeight='26px'>
              We can help.
            </Text>
          </Stack>
        </GridItem>
        <GridItem borderBottom='2px solid #11866f'>
          <NextLink href={`${DOCS_PAGE}/getting-started`} passHref>
            <Link _hover={{ textDecoration: 'none' }}>
              <Stack
                data-group
                bg='#d7f5ef'
                _hover={{ textDecoration: 'none', bg: '#11866f', color: '#f0f2e2' }}
                _focus={{
                  textDecoration: 'none',
                  bg: '#11866f',
                  color: '#f0f2e2',
                  boxShadow: 'inset 0 0 0 3px #f0f2e2 !important'
                }}
                _active={{ textDecoration: 'none', bg: '#25453f', color: '#f0f2e2' }}
                justifyContent='center'
                h='100%'
                p={4}
              >
                <Text
                  fontFamily='"JetBrains Mono", monospace'
                  // TODO: move to theme colors
                  fontWeight={700}
                  textTransform='uppercase'
                  textAlign='center'
                  color='#11866f'
                  _groupHover={{ color: '#f0f2e2' }}
                  _groupActive={{ color: '#f0f2e2' }}
                  _groupFocus={{ color: '#f0f2e2' }}
                >
                  Get started
                </Text>
              </Stack>
            </Link>
          </NextLink>
        </GridItem>

        {/* faq */}
        <GridItem borderRight='2px solid #11866f' borderBottom='2px solid #11866f'>
          <Stack p={4} h='100%'>
            <Text fontFamily='"Inter", sans-serif' lineHeight='26px' mb={-1}>
              Have doubts?
            </Text>

            <Text fontFamily='"Inter", sans-serif' lineHeight='26px'>
              Check the FAQ section in the documentation.
            </Text>
          </Stack>
        </GridItem>
        <GridItem borderBottom='2px solid #11866f'>
          <NextLink href={FAQ_PAGE} passHref>
            <Link _hover={{ textDecoration: 'none' }}>
              <Stack
                data-group
                bg='#d7f5ef'
                _hover={{ textDecoration: 'none', bg: '#11866F', color: '#f0f2e2' }}
                justifyContent='center'
                h='100%'
                p={4}
              >
                <Text
                  fontFamily='"JetBrains Mono", monospace'
                  // TODO: move to theme colors
                  fontWeight={700}
                  textTransform='uppercase'
                  textAlign='center'
                  color='#11866f'
                  _groupHover={{ color: '#f0f2e2' }}
                >
                  Go to the FAQ
                </Text>
              </Stack>
            </Link>
          </NextLink>
        </GridItem>

        {/* how to contribute */}
        <GridItem borderRight='2px solid #11866f'>
          <Stack p={4} h='100%'>
            <Text fontFamily='"Inter", sans-serif' lineHeight='26px' mb={-1}>
              Want to know how to contribute?
            </Text>

            <Text fontFamily='"Inter", sans-serif' lineHeight='26px'>
              Get more information in the documentation.
            </Text>
          </Stack>
        </GridItem>
        <GridItem>
          <NextLink href={CONTRIBUTING_PAGE} passHref>
            <Link _hover={{ textDecoration: 'none' }}>
              <Stack
                data-group
                bg='#d7f5ef'
                _hover={{ textDecoration: 'none', bg: '#11866F', color: '#f0f2e2' }}
                justifyContent='center'
                h='100%'
                p={4}
              >
                <Text
                  fontFamily='"JetBrains Mono", monospace'
                  // TODO: move to theme colors
                  fontWeight={700}
                  textTransform='uppercase'
                  textAlign='center'
                  color='#11866f'
                  _groupHover={{ color: '#f0f2e2' }}
                >
                  How to contribute
                </Text>
              </Stack>
            </Link>
          </NextLink>
        </GridItem>
      </Grid>
    </Stack>
  );
};
