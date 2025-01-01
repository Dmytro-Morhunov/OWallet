import React from 'react';
import {View, SafeAreaView, StyleSheet} from 'react-native';
import {useFormik} from 'formik';

import {colors} from '../utils';
import {Button, Input, Typography} from '../components';

export function AddWallet() {
  const form = useFormik({
    initialValues: {name: '', hash: ''},
    onSubmit: values => {
      console.log(values, 'form values');
    },
  });

  return (
    <View style={styles.container}>
      <SafeAreaView style={styles.safeArea}>
        <Typography text="Add Wallet" style={styles.title} />
        <Input
          placeholderText="Enter name"
          label="Name"
          value={form.values.name}
          onChangeText={form.handleChange('name')}
        />
        <Input
          placeholderText="Enter address"
          label="Hash address"
          value={form.values.hash}
          onChangeText={form.handleChange('hash')}
        />

        <View style={styles.buttons}>
          <Button title="Add wallet" onPress={form.handleSubmit} />
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
    paddingBottom: 24,
  },
  buttons: {
    flex: 1,
    justifyContent: 'flex-end',
  },
});
