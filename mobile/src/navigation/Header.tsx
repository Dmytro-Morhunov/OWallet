import React from 'react';
import {View, StyleSheet, Pressable} from 'react-native';
import {FontAwesomeIcon} from '@fortawesome/react-native-fontawesome';
import {faArrowLeft} from '@fortawesome/free-solid-svg-icons';
import {useSafeAreaInsets} from 'react-native-safe-area-context';
import {colors} from '../utils';
import {useNavigation} from '@react-navigation/native';

export function Header() {
  const navigation = useNavigation();
  const {top} = useSafeAreaInsets();

  const handleBackNavigation = () => {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  };

  return (
    <View style={[styles.container, {paddingTop: top}]}>
      <Pressable onPress={handleBackNavigation}>
        <View style={styles.button}>
          <FontAwesomeIcon
            size={20}
            icon={faArrowLeft}
            color={colors.secondary}
          />
        </View>
      </Pressable>
      <View>{null}</View>
      <View>{null}</View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    height: 100,
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'space-between',
    backgroundColor: colors.primary,
    alignItems: 'flex-end',
  },
  button: {
    alignSelf: 'center',
    padding: 16,
  },
});
