
import * as async from './async';

import {
    BRANCHES_LIST
} from '../actions/branches';


const initialState = {
    list: async.initialState,
};

export default  (state = initialState, action) => {
    state = {
        ...state,
        list: async.reduce(BRANCHES_LIST, state.list, action),
    };

    switch (action.type) {
        default:
            return state;
    }
};