import {createStaticNavigation} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import {Login} from '../screens/Login';
import {Home} from '../screens/Home';
import {Registration} from '../screens/Registration';
import {AddWallet} from '../screens/AddWallet';
import {Header} from './Header';

export enum Screens {
  HomeScreen = 'Home',
  LoginScreen = 'Login',
  RegistrationScreen = 'Registration',
  AddWalletScreen = 'AddCard',

  MainStack = 'Main',
  AuthStack = 'Auth',
}

export type HomeStackParamList = {
  [Screens.HomeScreen]: undefined;
  [Screens.AddWalletScreen]: undefined;
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
    [Screens.HomeScreen]: {screen: Home, options: {headerShown: false}},
    [Screens.AddWalletScreen]: {
      screen: AddWallet,
      options: {header: () => <Header />},
    },
  },
});

const AuthStack = createNativeStackNavigator<AuthStackParamList>({
  screens: {
    [Screens.LoginScreen]: {screen: Login, options: {headerShown: false}},
    [Screens.RegistrationScreen]: {
      screen: Registration,
      options: {header: () => <Header />},
    },
  },
});

const RootStack = createNativeStackNavigator<RootStackParamList>({
  screens: {[Screens.AuthStack]: AuthStack, [Screens.MainStack]: HomeStack},
  screenOptions: {
    headerShown: false,
  },
});

export const Navigation = createStaticNavigation(RootStack);
