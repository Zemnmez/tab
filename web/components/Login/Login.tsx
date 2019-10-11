import React from 'react';
import gql from 'graphql-tag';
import style from './Login.module.css';
import { Editor, EditorState } from 'draft-js';
import { useQuery } from '@apollo/react-hooks';
import raw from 'raw.macro';
const LOGIN_QUERY = gql(raw("./login.graphql"));

// TODO: apollo codegen

type Provider = {
    Name: string,
    Callback: string,
    AuthorizationEndpoint: string,
}

type loginQueryVars = {
    username: string
}

type loginQueryData = {
    User: {
        ByName: {
            Authentication: {
                OIDC: {
                    Providers: Array<Provider>
                }
            }
        }
    }
}

export const Login = ({ }) => {

    const [providers, setProviders] = 
        React.useState<Array<Provider>>();


    return <>
        <ValidatedField {...{
            placeholder: "username",
            useValue: (...args) => {
                const {loading, error, data} = useQuery<
                    loginQueryData,
                    loginQueryVars
                >(LOGIN_QUERY, ...args);

                return {
                    loading, error,
                    value: data, valid: !error,
                }
            },

            onChange: ({ value }: { value: loginQueryData }) => {
                const {
                    User: {
                        ByName: {
                            Authentication: {
                                OIDC: {
                                    Providers
                                }
                            }
                        }
                    }
                } = value;

                return setProviders(Providers);
            }
        }}/>

        {providers && <ProvidersList {...{
            providers
        }}/>}

    </>
}

export const ProvidersList: React.FunctionComponent<{
    providers: Array<Provider>
}> = ({ providers }) => <>
    <div {...{
        className: style.ProvidersList
    }}>
        {providers.map((provider, i) => <ProviderButton key={i} {...{provider}}/>) }
    </div>
</>

export const ProviderButton: React.FunctionComponent<{
    provider: Provider
}> = ({ provider }) => <>
    <div {...{
        className: style.ProviderButton
    }}>
        <a {...{
            href: provider.AuthorizationEndpoint,
        }}>{provider.Name}</a>
    </div>
</>

export const ValidatedField = ({ useValue, onChange, placeholder }) => {
    const [editorState, setEditorState ] =
        React.useState(EditorState.createEmpty());

    const { loading, error, value, valid } = useValue(
        editorState.getCurrentContent().getPlainText()
    );

    React.useEffect(() => { onChange({ error, value}) }, [error, value]);

    return <Editor {...{
        className: [
            style.ValidatedField,
            valid? style.valid: '',
            loading? style.loading: '',
            error? style.error: '',
        ].join(" "),

        editorState,

        onChange: setEditorState,
    }}/>
}