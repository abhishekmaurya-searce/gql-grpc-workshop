import React from "react";
import { useState } from "react";
import { GREET_MUTATION } from "../graphql/Mutation";
import { useMutation } from "@apollo/client";
function GreetForm() {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
  
    const [greet] = useMutation(GREET_MUTATION);
  
    const greeting = () => {
      greet({
        variables: {
          input:{
            firstName: firstName,
            lastName: lastName,
          }
        },
      });

    };
    return (
      <div>
        <input
          type="text"
          placeholder="First Name"
          onChange={(e) => {
            setFirstName(e.target.value);
          }}
        />
        <input
          type="text"
          placeholder="Last Name"
          onChange={(e) => {
            setLastName(e.target.value);
          }}
        />
        <button onClick={greeting}>Greet</button>
        {/* {greet && <p style={{ color: 'red' }}>{greet}</p>} */}
      </div>
    );
  }
  export default GreetForm;