import {createStaticNavigation} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import {Login} from '../screens/Login';
import {Home} from '../screens/Home';
import {Registration} from '../screens/Registration';

export enum Screens {
  HomeScreen = 'Home',
  LoginScreen = 'Login',
  RegistrationScreen = 'Registration',
  MainStack = 'Main',
  AuthStack = 'Auth',
}

export type HomeStackParamList = {
  [Screens.HomeScreen]: undefined;
};

export type AuthStackParamList = {
  [Screens.LoginScreen]: undefined;
  [Screens.RegistrationScreen]: undefined;
};

export type RootStackParamList = {
  [Screens.AuthStack]: undefined;
  [Screens.MainStack]: undefined;
};

const HomeStack = createNativeStackNavigator<HomeStackParamList>({
  screens: {
    [Screens.HomeScreen]: Home,
  },
});

const AuthStack = createNativeStackNavigator<AuthStackParamList>({
  screens: {
    [Screens.LoginScreen]: Login,
    [Screens.RegistrationScreen]: Registration,
  },
});

const RootStack = createNativeStackNavigator<RootStackParamList>({
  screens: {[Screens.AuthStack]: AuthStack, [Screens.MainStack]: HomeStack},
  screenOptions: {
    headerShown: false,
  },
});

export const Navigation = createStaticNavigation(RootStack);
