import React from 'react';
import {
  View,
  SafeAreaView,
  StyleSheet,
  FlatList,
  ListRenderItem,
} from 'react-native';
import {NavigationProp, useNavigation} from '@react-navigation/native';
import {useSafeAreaInsets} from 'react-native-safe-area-context';

import {colors} from '../utils';
import {Card, Typography, Button} from '../components';
import {HomeStackParamList, Screens} from '../navigation';

interface Wallet {
  name: string;
  address: string;
  balance: string;
}

const mockWallets: Wallet[] = [
  {name: 'Bitcoin', address: 'Some address', balance: '1'},
  {name: 'Ethereum', address: 'Some address 2', balance: '0.2341'},
  {name: 'DOGE', address: 'Some address 3', balance: '123123123'},
];

export function Home() {
  const {top} = useSafeAreaInsets();
  const navigation = useNavigation<NavigationProp<HomeStackParamList>>();

  const navigateToAddWallet = () => {
    navigation.navigate(Screens.AddWalletScreen);
  };
  const handleCard = () => {
    console.log('handle card');
  };

  const _renderCard: ListRenderItem<Wallet> = ({item}) => (
    <Card
      name={item.name}
      address={item.address}
      balance={item.balance}
      containerStyle={styles.card}
      onPress={handleCard}
    />
  );

  return (
    <View style={[styles.container, {paddingTop: top + 24}]}>
      <SafeAreaView style={styles.safeArea}>
        <Typography text="Wallets" />
        <FlatList data={mockWallets} renderItem={_renderCard} />
        <Button title="Add wallet" onPress={navigateToAddWallet} />
      </SafeAreaView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 24,
    backgroundColor: colors.primary,
  },
  safeArea: {
    flex: 1,
  },
  card: {
    marginBottom: 20,
  },
});
