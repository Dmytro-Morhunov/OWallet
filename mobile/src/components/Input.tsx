import React from 'react';
import {
  View,
  Text,
  TextInput,
  StyleSheet,
  StyleProp,
  ViewStyle,
} from 'react-native';
import {colors, fonts} from '../utils';

interface Props {
  label: string;
  value: string;
  error?: string;
  placeholderText: string;
  containerStyle?: StyleProp<ViewStyle>;

  onChangeText: (str: string) => void;
}

export function Input({
  label,
  error,
  value,
  placeholderText,
  containerStyle,
  onChangeText,
}: Props) {
  return (
    <View style={containerStyle}>
      <Text style={styles.label}>{label}</Text>
      <TextInput
        placeholder={placeholderText}
        style={styles.input}
        value={value}
        onChangeText={onChangeText}
      />
      <Text style={styles.errorMessage}>{error ? error : null}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  label: {
    color: colors.white,
    fontFamily: fonts.RockwellRegular,
    fontSize: 20,
    lineHeight: 22,
    paddingBottom: 6,
  },
  input: {
    color: colors.black,
    fontFamily: fonts.RockwellRegular,
    fontSize: 20,
    lineHeight: 22,
    backgroundColor: colors.white,
    height: 60,
    paddingHorizontal: 12,
  },
  errorMessage: {
    fontSize: 12,
    fontFamily: fonts.RockwellRegular,
    color: colors.red,
    lineHeight: 14,
    paddingTop: 6,
  },
});
