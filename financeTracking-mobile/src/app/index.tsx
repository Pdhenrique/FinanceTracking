import Login from '@/screen/login'
import { View, Text, Button } from 'react-native'

export default function Index() {

    return (
        <View style={{ justifyContent: 'center', alignItems: "center", flex: 1 }}>
            <Login />
        </View>
    )
}