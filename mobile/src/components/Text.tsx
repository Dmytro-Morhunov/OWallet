import React from 'react';
import {Text, StyleSheet, StyleProp, TextStyle} from 'react-native';
import {colors, fonts} from '../utils';

interface Props {
  text: string;
  style?: StyleProp<TextStyle>;
}

export function Typography({text, style}: Props) {
  return <Text style={[styles.text, style]}>{text}</Text>;
}

const styles = StyleSheet.create({
  text: {
    fontSize: 40,
    lineHeight: 46,
    fontFamily: fonts.RockwellRegular,
    color: colors.white,
  },
});
