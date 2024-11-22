import React from 'react';
import {
  View,
  Text,
  TextInput,
  StyleSheet,
  TouchableOpacity,
} from 'react-native';
import {useFormik} from 'formik';
import {AuthStackParamList, Screens} from '../navigation';
import {NavigationProp, useNavigation} from '@react-navigation/native';

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
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Email</Text>
      <TextInput
        style={styles.input}
        onChangeText={form.handleChange('email')}
        onBlur={form.handleBlur('email')}
        value={form.values.email}
      />
      <Text style={styles.title}>Password</Text>
      <TextInput
        style={styles.input}
        onChangeText={form.handleChange('password')}
        onBlur={form.handleBlur('password')}
        value={form.values.password}
      />

      <TouchableOpacity style={styles.button} onPress={navigateToMain}>
        <Text>Submit</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {flex: 1, justifyContent: 'center', paddingHorizontal: 32},
  title: {
    fontSize: 12,
    paddingVertical: 8,
  },
  input: {
    width: '100%',
    height: 42,
    borderWidth: 1,
    borderColor: 'black',
  },
  button: {
    marginTop: 16,
    width: '100%',
    height: 42,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'gray',
  },
});
