// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.
import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import Unity, { UnityContent } from 'react-unity-webgl';

import { WebSocketClient } from '@/api/websockets';
import { RouteConfig } from '@/app/routes';
import { getMatchScore } from '@/app/store/actions/mathes';
import { ToastNotifications } from '@/notifications/service';
import { getCurrentWebSocketClient } from '@/webSockets/service';

import './index.scss';

/** Variable describes that game is over. */
const MATCH_RESULT: string = 'do you allow us to take your address?';

const FootballGame: React.FC = () => {
    const history = useHistory();
    const dispatch = useDispatch();

    const [webSocketClient, setWebSocketClient] = useState<WebSocketClient | null>(null);

    const unityContext = new UnityContent('/static/dist/webGl/football.json', '/static/dist/webGl/UnityLoader.js');

    if (webSocketClient) {
        webSocketClient.ws.onmessage = ({ data }: MessageEvent) => {
            const event = JSON.parse(data);

            if (event.message.question === MATCH_RESULT) {
                dispatch(getMatchScore(event.message));
                history.push(RouteConfig.Match.path);
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

    return <Unity unityContent={unityContext} className="unity-container" />;
};

export default FootballGame;
