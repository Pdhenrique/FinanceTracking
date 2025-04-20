import { View, Text, Button, TextInput } from 'react-native'

function Login() {
    return (
        <View>
            <TextInput id="login-field-email" placeholder='Email'></TextInput>

            <TextInput id="login-field-password" placeholder='password'></TextInput>
        </View>
    )
}

export default Login