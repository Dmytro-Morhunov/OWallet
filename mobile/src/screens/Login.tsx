import React from 'react';
import {View, StyleSheet, SafeAreaView} from 'react-native';
import {useFormik} from 'formik';
import {NavigationProp, useNavigation} from '@react-navigation/native';


import {AuthStackParamList, Screens} from '../navigation';
import {Button, ButtonType, Input, Typography} from '../components';
import {colors} from '../utils';

export function Login() {
  const navigation = useNavigation<NavigationProp<AuthStackParamList>>();
  const form = useFormik({
    initialValues: {email: '', password: ''},
    onSubmit: values => {
      console.log(values, 'form values');
    },
  });

  const navigateToMain = () => {
    navigation.navigate(Screens.MainStack);
  };

  const navigateToRegistration = () => {
    navigation.navigate(Screens.RegistrationScreen);
  };

  return (
    <View style={styles.container}>
      <SafeAreaView style={styles.safeArea}>
        <Typography text="Login" style={styles.title} />
        <Input
          label="Email"
          placeholderText="Enter your email"
          value={form.values.email}
          error={form.errors.email}
          onChangeText={form.handleChange('email')}
        />
        <Input
          label="Password"
          placeholderText="Enter your password"
          value={form.values.password}
          error={form.errors.password}
          onChangeText={form.handleChange('password')}
        />

        <View style={styles.buttons}>
          <Button
            title="Login"
            containerStyle={styles.loginButton}
            onPress={navigateToMain}
          />
          <Button
            title="Create account"
            type={ButtonType.secondary}
            onPress={navigateToRegistration}
          />
        </View>
      </SafeAreaView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingHorizontal: 24,
    backgroundColor: colors.primary,
  },
  safeArea: {
    flex: 1,
  },
  title: {
    paddingVertical: 32,
  },
  loginButton: {
    marginBottom: 18,
  },
  buttons: {
    flex: 1,
    justifyContent: 'flex-end',
  },
});
