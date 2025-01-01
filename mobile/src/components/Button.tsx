import React, {useMemo} from 'react';
import {
  Pressable,
  Text,
  StyleSheet,
  View,
  StyleProp,
  ViewStyle,
} from 'react-native';
import {colors, fonts} from '../utils';

export enum ButtonType {
  primary = 'primary',
  secondary = 'secondary',
}

interface Props {
  title: string;
  type?: ButtonType;
  containerStyle?: StyleProp<ViewStyle>;
  onPress: () => void;
}

export function Button({
  title,
  type = ButtonType.primary,
  onPress,
  containerStyle,
}: Props) {
  const style = useMemo(() => {
    const defaultStyle: StyleProp<ViewStyle> = {};
    switch (type) {
      case ButtonType.secondary: {
        defaultStyle.backgroundColor = colors.gray;
      }
    }
    return defaultStyle;
  }, [type]);

  return (
    <Pressable onPress={onPress}>
      <View style={[styles.container, containerStyle, style]}>
        <Text style={styles.text}>{title}</Text>
      </View>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    height: 54,
    backgroundColor: colors.secondary,
    justifyContent: 'center',
    alignItems: 'center',
  },
  text: {
    fontSize: 20,
    lineHeight: 22,
    fontFamily: fonts.RockwellRegular,
  },
});
