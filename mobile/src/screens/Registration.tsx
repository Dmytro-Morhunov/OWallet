import React from 'react';
import {View, SafeAreaView, StyleSheet} from 'react-native';
import {useFormik} from 'formik';

import {Typography, Input, Button} from '../components';
import {colors} from '../utils';

export function Registration() {
  const form = useFormik({
    initialValues: {email: '', password: ''},
    onSubmit: values => {
      console.log(values, 'form values');
    },
  });

  return (
    <View style={styles.container}>
      <SafeAreaView style={styles.safeArea}>
        <Typography text="Create new account" style={styles.title} />
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
        <Input
          label="Confirm password"
          placeholderText="Confirm your password"
          value={form.values.password}
          error={form.errors.password}
          onChangeText={form.handleChange('password')}
        />

        <View style={styles.buttons}>
          <Button title="Create" onPress={() => console.log('Registration')} />
        </View>
      </SafeAreaView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.primary,
    padding: 24,
  },
  safeArea: {flex: 1},
  title: {
    paddingBottom: 32,
  },
  buttons: {
    flex: 1,
    justifyContent: 'flex-end',
  },
});
