import {ApolloClient, InMemoryCache,ApolloProvider} from '@apollo/client'
import GreetForm from './components/GreetForm';
import Greet from './components/Greet'
const client = new ApolloClient({
  uri: 'http://localhost:8080/greet',
  cache: new InMemoryCache(),
});
function App() {
  
  return (
    <ApolloProvider client={client}>
      <div><h1>Apollo</h1>
      <GreetForm></GreetForm>
      <h1>Http</h1>
      <Greet/>
      </div>
    </ApolloProvider>
  );
}


export default App;
