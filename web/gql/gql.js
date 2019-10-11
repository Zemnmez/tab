import ApolloClient from 'apollo-boost';

const mustEnv = (name) => {
    const ret = process.env[name];
    if (!ret) throw new Error(`missing env var ${name}`);

    return ret;
}

export const gqlClient = new ApolloClient({
    uri: mustEnv("GQL_SERVER")
})

