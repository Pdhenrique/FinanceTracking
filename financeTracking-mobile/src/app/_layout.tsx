import { Stack } from "expo-router";
import { Text, View } from "react-native";
import { colors } from "@/styles/theme";

import {
    useFonts,
    Rubik_600SemiBold,
    Rubik_400Regular,
    Rubik_700Bold,
    Rubik_500Medium,
} from "@expo-google-fonts/rubik";

export default function Layout() {
    const [fontsLoaded] = useFonts({
        Rubik_600SemiBold,
        Rubik_400Regular,
        Rubik_700Bold,
        Rubik_500Medium,
    });

    if (!fontsLoaded) {
        return
    }

    return (
        <Stack
            screenOptions={{
                headerShown: false,
                contentStyle: { backgroundColor: colors.gray[100] },
            }}
        />
    );
}
