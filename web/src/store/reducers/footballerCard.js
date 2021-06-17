const card = {
    overalInfo: {
        'Name': 'Albert Ronalculus',
        'Nation': 'Portugal 🇵🇹',
        'Skills': '5',
        'Weak foot': '4',
        'Intl. Rep': '5',
        'Foot': 'Right',
        'Height': '187',
        'Nation?': '83',
        'Revision': 'Rare',
        'Def. WR': 'Low',
        'Att. WR': 'High',
        'Added on': '2020-09-10',
        'Origin': 'NA',
        'R.Face': 'Low',
        'B.Type': true,
        'Age': '36',
    },
    tactics: {
        'Tactics': 98,
        'Positioning': 70,
        'Composure': 70,
        'Aggression': 70,
        'Vision': 70,
        'Awareness': 70,
        'Crosses': 70,
    },
    physique: {
        'Physique': 34,
        'Acceleration': 70,
        'Running speed': 70,
        'Reaction speed': 70,
        'Agility': 70,
        'Stamina': 70,
        'Strength': 70,
        'Jumping': 70,
        'Balance': 70,
    },
    technique: {
        'Technique': 26,
        'Dribbing': 70,
        'Ball Control': 70,
        'Weak Foot': 70,
        'Skill Moves': 70,
        'Finesse': 70,
        'Curve': 70,
        'Volleys': 70,
        'Short passing': 70,
        'Long passing': 70,
        'Forward pass': 70,
    },
    offence: {
        'Offence': 42,
        'Finishing ability': 70,
        'Shot power': 70,
        'Accuracy': 70,
        'Distance': 70,
        'Penalty': 70,
        'Free Kicks': 70,
        'Corners': 70,
        'Heading accuracy': 70,
    },
    defence: {
        'Defence': 74,
        'Offside trap': 70,
        'Tackles': 70,
        'Ball focus': 70,
        'Interceptions': 70,
        'Vigilance': 70,
    },
    goalkeeping: {
        'Goalkeeping': 84,
        'Diving': 70,
        'Handling': 70,
        'Sweeping': 70,
        'Throwing': 70,
    }
};

function cardlist(count) {
    let quantity = count;
    const list = [];

    while (quantity > 0) {
        list.push(card);
        quantity--;
    }

    return list;
}
/* eslint-disable */
export const cardReducer = (cardState = cardlist(15), action) => {
    return cardState;
};