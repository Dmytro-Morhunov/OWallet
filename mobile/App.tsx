import React, {useEffect} from 'react';
import {Navigation} from './src/navigation/index.tsx';
import {Provider} from 'react-redux';
import BootSplash from 'react-native-bootsplash';

import {store} from './src/data/store/createStore.ts';

function App(): React.JSX.Element {
  useEffect(() => {
    const init = async () => {
      // …do multiple sync or async tasks
    };

    init().finally(async () => {
      await BootSplash.hide({fade: true});
      console.log('BootSplash has been hidden successfully');
    });
  }, []);

  return (
    <Provider store={store}>
      <Navigation
        onReady={() => {
          BootSplash.hide();
        }}
      />
    </Provider>
  );
}

export default App;
