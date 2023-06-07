// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';

import { AutoCloseTimer } from './AutoCloseTimer';
import { Timer } from './Timer';

import { WebSocketClient } from '@/api/websockets';
import { RouteConfig } from '@/app/routes';
import { RootState } from '@/app/store';
import { getMatchScore } from '@/app/store/actions/mathes';
import { startSearchingMatch } from '@/app/store/actions/clubs';
import { getCurrentWebSocketClient, onCloseConnection, onOpenConnectionNoAction, sendAction, setMatchQueue } from '@/webSockets/service';
import { ToastNotifications } from '@/notifications/service';

import './index.scss';

const MatchFinder: React.FC = () => {
    const { squad } = useSelector(
        (state: RootState) => state.clubsReducer.activeClub
    );
    const { isSearchingMatch } = useSelector(
        (state: RootState) => state.clubsReducer
    );

    const [webSocketClient, setWebSocketClient] = useState<WebSocketClient | null>(null);

    const dispatch = useDispatch();
    const history = useHistory();

    /** Indicates if match is found. */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** Indicates if match is confirmed. */
    const [isMatchConfirmed, setIsMatchConfirmed] = useState<boolean>(false);

    /** CANCEL_GAME_DELAY_MIN is min time delay for auto cancel game. */
    const CANCEL_GAME_DELAY_MIN: number = 29000;
    /** CANCEL_GAME_DELAY_MAX is max time delay for auto cancel game. */
    const CANCEL_GAME_DELAY_MAX: number = 31000;

    // TODO: it will be deleted after ./gameplage/queue/chore.go solution.
    /** Returns random time delay from range for auto cancel game. */
    function getRandomTimeDelayForCancelGame(minTimeDelay: number, maxTimeDelay: number) {
        return Math.random() * (maxTimeDelay - minTimeDelay) + minTimeDelay;
    };

    /** Delay is time delay for redirect user to match page. */
    const DELAY: number = 2000;

    /** Variable describes that webscoket connection responsed with error. */
    const ERROR_MESSAGE: string = 'could not write to websocket';
    /** Variable describes that user still searching game. */
    const STILL_SEARCHING_MESSAGE: string = 'you are still in search';
    /** Variable describes that was send wrong action from user. */
    const WRONG_ACTION_MESSAGE: string = 'wrong action';
    /** Variable describes that user added to gueue. */
    const YOU_ADDED_MESSAGE: string = 'you added';
    /** Variable describes that it needs confirm game from user. */
    const YOU_CONFIRM_PLAY_MESSAGE: string = 'do you confirm play?';
    /** Variable describes that user have leaved from searching game. */
    const YOU_LEAVED_MESSAGE: string = 'you left';
    /** Variable describes that user have leaved from searching game. */
    const PLAYERS_FOUND: string = 'players found';

    /** Sends confirm action. */
    const confirmMatch = () => {
        setIsMatchConfirmed(true);
        sendAction('confirm', squad.id);
    };

    /** Canceles confirmation game. */
    const cancelConfirmationGame = () => {
        sendAction('reject', squad.id);
    };

    /** Canceles searching game and closes MatchFinder component. */
    const canselSearchingGame = () => {
        setMatchQueue();

        sendAction('finishSearch', squad.id);

        /** Updates current queue client. */
        const updatedClient = getCurrentWebSocketClient();
        setWebSocketClient(updatedClient);

        dispatch(startSearchingMatch(false));
    };

    /** Exposes start searching match logic. */
    const startSearchMatch = () => {
        setMatchQueue();

        sendAction('startSearch', squad.id);

        /** Updates current websocket client. */
        const newclient = getCurrentWebSocketClient();
        setWebSocketClient(newclient);
    };

    useEffect(() => {
        isSearchingMatch && startSearchMatch();
    }, [isSearchingMatch]);

    /** Processes websocket client event messages. */
    if (webSocketClient) {
        webSocketClient.ws.onmessage = ({ data }: MessageEvent) => {
            const event = JSON.parse(data);

            switch (event.message) {
            case ERROR_MESSAGE:
                ToastNotifications.notify('error message');

                return;
            case STILL_SEARCHING_MESSAGE:
                ToastNotifications.gameCanceled();

                return;
            case YOU_ADDED_MESSAGE:
                setIsMatchFound(false);

                return;
            case YOU_CONFIRM_PLAY_MESSAGE:
                setIsMatchFound(true);
                setIsMatchConfirmed(false);

                return;

            case YOU_LEAVED_MESSAGE:
                setIsMatchFound(false);
                dispatch(startSearchingMatch(false));

                onCloseConnection();
                onOpenConnectionNoAction();

                return;

            case PLAYERS_FOUND:
                setIsMatchFound(false);
                ToastNotifications.gameFinished();

                dispatch(startSearchingMatch(false));

                /** implements redirect to game page after DELAY time.  */
                setTimeout(() => {
                    history.push(RouteConfig.FootballGame.path);
                }, DELAY);
                break;
            default:
                /** no actions */
            }
        };
    }

    if (webSocketClient) {
        webSocketClient.ws.onerror = (event: Event) => {
            ToastNotifications.somethingWentsWrong();
        };
    }

    useEffect(() => {
        /** Canceles confirm game after CANCEL_GAME_DELAY delay. */
        let autoCancelConfirmGame: ReturnType<typeof setTimeout>;
        if (isMatchFound) {
            autoCancelConfirmGame = setTimeout(() => {
                sendAction('reject', squad.id);
            }, getRandomTimeDelayForCancelGame(CANCEL_GAME_DELAY_MIN, CANCEL_GAME_DELAY_MAX));
        }

        return () => clearTimeout(autoCancelConfirmGame);
    }, [isMatchFound]);

    return isSearchingMatch && <section className={isMatchFound ? 'match-finder__wrapper' : ''}>
        <div className="match-finder">
            <h1 className="match-finder__title">
                {isMatchFound ? 'YOUR MATCH WAS FOUND' : 'LOOKING FOR A MATCH'}
            </h1>
            {isMatchFound ? <AutoCloseTimer /> : <Timer />}
            <div className="match-finder__form">
                {isMatchFound ? <input
                    className="match-finder__form__accept"
                    value="Accept"
                    type="button"
                    onClick={confirmMatch}
                /> : <input
                    className="match-finder__form__accept-not-allowed"
                    value="Accept"
                    type="button"
                />}
                {isMatchFound ? <input
                    className={`match-finder__form__cancel${isMatchConfirmed ? '-not-allowed' : ''}`}
                    value="Cancel"
                    type="button"
                    onClick={cancelConfirmationGame}
                /> : <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={canselSearchingGame}
                />}
            </div>
        </div>
    </section>;
};

export default MatchFinder;
