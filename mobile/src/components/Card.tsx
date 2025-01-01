import React from 'react';
import {View, Pressable, StyleSheet, StyleProp, ViewStyle} from 'react-native';

import {Typography} from './Text';
import {colors, fonts} from '../utils';

interface Props {
  name: string;
  address: string;
  balance: string;
  containerStyle?: StyleProp<ViewStyle>;

  onPress: () => void;
}

export function Card({name, address, balance, containerStyle, onPress}: Props) {
  return (
    <Pressable onPress={onPress}>
      <View style={[styles.container, containerStyle]}>
        <Typography style={styles.mediumText} text={name} />
        <Typography style={styles.mediumText} text={address} />
        <Typography style={styles.balance} text={balance} />
      </View>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    backgroundColor: colors.secondary,
    padding: 16,
    borderRadius: 8,
  },
  mediumText: {
    fontSize: 20,
    lineHeight: 22,
    fontFamily: fonts.RockwellRegular,
    paddingBottom: 16,
    color: colors.black,
  },
  balance: {
    fontSize: 50,
    lineHeight: 56,
    fontFamily: fonts.RockwellRegular,
    color: colors.black,
  },
});
