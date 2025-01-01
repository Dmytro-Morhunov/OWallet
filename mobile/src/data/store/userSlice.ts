import {createSlice} from '@reduxjs/toolkit';
import type {PayloadAction} from '@reduxjs/toolkit';

export interface UserState {
  token: string | null;
}

const initialState: UserState = {
  token: null,
};

export const counterSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
});

// Action creators are generated for each case reducer function
// export const {increment, decrement, incrementByAmount} = counterSlice.actions;

export default counterSlice.reducer;
