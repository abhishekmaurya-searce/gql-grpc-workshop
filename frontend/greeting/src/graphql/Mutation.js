import { gql } from "@apollo/client";

export const GREET_MUTATION  = gql`
    mutation greet($input: UserInput!) {
              greet(input: $input) {
                result
              }
            }

`