import React, { useState } from 'react';

function Greet() {
  const [firstname, setFirstname] = useState('');
  const [lastname, setLastname] = useState('');
  const [error, setError] = useState(null);
  const [msg,setMsg] = useState('');
  const handleGreet = async () => {
    try {
      const response = await fetch('http://localhost:8080/greet', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          query: `
            mutation Greet($input: UserInput!) {
              greet(input: $input) {
                result
              }
            }
          `,
          variables: {
            input: {
              firstname,
              lastname,
            },
          },
        }),
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      console.log(data.data.greet.result);
      setMsg(data.data.greet.result)
      setError(null); // Clear any previous error
    } catch (error) {
      console.error('Network request error:', error);
      setError('An error occurred while sending the request.');
    }
  };

  return (
    <div>
      <input
        type="text"
        placeholder="First Name"
        value={firstname}
        onChange={(e) => setFirstname(e.target.value)}
      />
      <input
        type="text"
        placeholder="Last Name"
        value={lastname}
        onChange={(e) => setLastname(e.target.value)}
      />
      <button onClick={handleGreet}>Greet</button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {msg && <p style={{ color: 'red' }}>{msg}</p>}
    </div>
  );
}

export default Greet;
