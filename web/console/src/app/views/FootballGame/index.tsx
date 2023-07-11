// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.
import { useCallback, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { Unity, useUnityContext } from 'react-unity-webgl';

import { WebSocketClient } from '@/api/websockets';
import { getCurrentWebSocketClient, sendUnityAction } from '@/webSockets/service';
import { ToastNotifications } from '@/notifications/service';
import { RouteConfig } from '@/app/routes';
import { getMatchScore } from '@/app/store/actions/mathes';

import './index.scss';

/** Describes that game is over. */
const MATCH_RESULT: string = 'do you allow us to take your address?';

/** Describes unity action to get start info for game. */
const START_UNITY_ACTION: string = 'GoodBye';

/** Describes unity action to get info for game. */
const UNITY_ACTION: string = 'PlayerAction';

/** Describes action to send game info in WS. */
const GAME_INFO_ACTION: string = 'send unity game information';

/** Describes action to send start game info in WS. */
const START_GAME_INFO_ACTION: string = 'send start unity game information';

/** Describes message of getting start game info. */
const GAME_START_INFO_MESSAGE: string = 'football information';

/** Describes message of getting game info. */
const GAME_INFO_MESSAGE: string = 'match action';

/** Describes game object name in unity to send message. */
const UNITY_GAME_OBJECT_NAME: string = 'Connection';

/** Describes unity method name. */
const UNITY_OBJECT_METHOD_NAME: string = 'ReciveActionMessage';

/** Describes unity method name. */
const START_UNITY_OBJECT_METHOD_NAME: string = 'ReciveBaseMessage';

const FootballGame: React.FC = () => {
    const history = useHistory();
    const dispatch = useDispatch();

    const [webSocketClient, setWebSocketClient] = useState<WebSocketClient | null>(null);

    const { sendMessage, unityProvider, addEventListener, removeEventListener } = useUnityContext({
        loaderUrl: '/static/dist/webGl/Football.loader.js',
        dataUrl: '/static/dist/webGl/Football.data',
        frameworkUrl: '/static/dist/webGl/Football.framework.js',
        codeUrl: '/static/dist/webGl/Football.wasm',
    });

    const handleStartUnityAction = useCallback((message) => {
        sendUnityAction(START_GAME_INFO_ACTION, message);
    }, []);

    const handleUnityActions = useCallback((message) => {
        sendUnityAction(GAME_INFO_ACTION, message);
    }, []);

    if (webSocketClient) {
        webSocketClient.ws.onmessage = ({ data }: MessageEvent) => {
            const event = JSON.parse(data);

            if (event.message === MATCH_RESULT) {
                dispatch(getMatchScore(event.message));
                // history.push(RouteConfig.Match.path);
            }

            switch (event.message) {
            case GAME_START_INFO_MESSAGE:
                sendMessage(UNITY_GAME_OBJECT_NAME, START_UNITY_OBJECT_METHOD_NAME, JSON.stringify(event.gameInformation));
                break;
            case GAME_INFO_MESSAGE:
                sendMessage(UNITY_GAME_OBJECT_NAME, UNITY_OBJECT_METHOD_NAME, JSON.stringify(event.gameInformation));
                break;
            default:
            }
        };
    }

    if (webSocketClient) {
        webSocketClient.ws.onerror = (event: Event) => {
            ToastNotifications.somethingWentsWrong();
        };
    }

    useEffect(() => {
        /** Updates current websocket client. */
        const newclient = getCurrentWebSocketClient();
        setWebSocketClient(newclient);
    }, []);

    useEffect(() => {
        addEventListener(START_UNITY_ACTION, handleStartUnityAction);
        addEventListener(UNITY_ACTION, handleUnityActions);

        return () => {
            removeEventListener(START_UNITY_ACTION, handleStartUnityAction);
            removeEventListener(UNITY_ACTION, handleUnityActions);
        };
    }, [addEventListener, removeEventListener, handleUnityActions, handleStartUnityAction]);

    return <Unity unityProvider={unityProvider} className="unity-container" />;
};
export default FootballGame;
