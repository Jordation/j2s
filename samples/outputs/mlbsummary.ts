interface ship {
  combined: object[];
}

interface assists {
  total: number;
  outfield: number;
}

interface away {
  statistics: statistics;
  remaining: number;
  loss: number;
  lineup: object[];
  players: object[];
  market: string;
  roster: object[];
  used: number;
  runs: number;
  errors: number;
  abbr: string;
  name: string;
  id: string;
  win: number;
  probablePitcher: probablePitcher;
  startingPitcher: startingPitcher;
  hits: number;
  scoring: object[];
}

interface broadcasts {
  network: string;
  type: string;
  locale: string;
  channel: string;
}

interface bullpen {
  oab: number;
  inPlay: inPlay;
  lob: number;
  era: number;
  outcome: outcome;
  outs: outs;
  k9: number;
  onbase: onbase;
  whip: number;
  timesThroughOrder: number;
  games: games;
  oba: number;
  gofo: number;
  gbfb: number;
  bk: number;
  bfIp: number;
  slg: number;
  obp: number;
  ip2: number;
  bf: number;
  pitches: pitches;
  wp: number;
  pitchCount: number;
  steal: steal;
  ip1: number;
  runs: runs;
  babip: number;
  kbb: number;
}

interface combined {
  game: game;
  Comment: string;
}

interface currentConditions {
  condition: string;
  humidity: number;
  dewPointF: number;
  cloudCover: number;
  obsTime: string;
  wind: wind;
  tempF: number;
}

interface errors {
  interference: number;
  total: number;
  throwing: number;
  fielding: number;
}

interface fielding {
  overall: overall;
  positions: object[];
}

interface final {
  inning: number;
  inningHalf: string;
}

interface forecast {
  dewPointF: number;
  cloudCover: number;
  obsTime: string;
  wind: wind;
  tempF: number;
  condition: string;
  humidity: number;
}

interface game {
  awayTeam: string;
  reviews: reviews;
  id: string;
  homeTeam: string;
  duration: string;
  entryMode: string;
  seasonId: string;
  pitching: pitching;
  seasonYear: number;
  timeZones: timeZones;
  gameNumber: number;
  dayNight: string;
  attendance: number;
  weather: weather;
  scheduled: string;
  moundVisits: moundVisits;
  doubleHeader: boolean;
  broadcasts: object[];
  home: home;
  status: string;
  coverage: string;
  venue: venue;
  officials: object[];
  reference: string;
  final: final;
  away: away;
  seasonType: string;
}

interface games {
  complete: number;
  hold: number;
  teamWin: number;
  start: number;
  save: number;
  blownSave: number;
  play: number;
  finish: number;
  teamLoss: number;
  svo: number;
  win: number;
  loss: number;
  qstart: number;
  shutout: number;
}

interface hitting {
  overall: overall;
}

interface hold {
  blownSave: number;
  lastName: string;
  jerseyNumber: string;
  position: string;
  primaryPosition: string;
  win: number;
  hold: number;
  fullName: string;
  preferredName: string;
  firstName: string;
  status: string;
  id: string;
  loss: number;
  save: number;
}

interface home {
  remaining: number;
  market: string;
  probablePitcher: probablePitcher;
  roster: object[];
  name: string;
  statistics: statistics;
  abbr: string;
  errors: number;
  lineup: object[];
  runs: number;
  loss: number;
  win: number;
  hits: number;
  startingPitcher: startingPitcher;
  id: string;
  used: number;
  scoring: object[];
  players: object[];
}

interface inPlay {
  linedrive: number;
  groundball: number;
  popup: number;
  flyball: number;
}

interface lineup {
  order: number;
  position: number;
  sequence: number;
  id: string;
  inningHalf: string;
  inning: number;
}

interface location {
  lat: string;
  lng: string;
}

interface loss {
  primaryPosition: string;
  loss: number;
  fullName: string;
  lastName: string;
  status: string;
  id: string;
  win: number;
  hold: number;
  preferredName: string;
  firstName: string;
  position: string;
  jerseyNumber: string;
  save: number;
  blownSave: number;
}

interface moundVisits {
  home: home;
  away: away;
}

interface officials {
  fullName: string;
  firstName: string;
  lastName: string;
  assignment: string;
  experience: string;
  id: string;
}

interface onbase {
  fc: number;
  hr9: number;
  hr: number;
  bb: number;
  rov: number;
  hbp: number;
  h: number;
  t: number;
  cycle: number;
  ci: number;
  h9: number;
  s: number;
  d: number;
  ibb: number;
  roe: number;
  tb: number;
}

interface outcome {
  ball: number;
  iball: number;
  dirtball: number;
  foul: number;
  klook: number;
  kswing: number;
  ktotal: number;
}

interface outs {
  fo: number;
  lo: number;
  lidp: number;
  ktotal: number;
  gidp: number;
  sacfly: number;
  po: number;
  fidp: number;
  klook: number;
  sachit: number;
  go: number;
  kswing: number;
}

interface overall {
  rf: number;
  bbk: number;
  gofo: number;
  fpct: number;
  bip: number;
  flyball: number;
  rbi: number;
  ap: number;
  a: number;
  abk: number;
  rbi2out: number;
  outs: outs;
  ab: number;
  dp: number;
  groundball: number;
  timesThroughOrder: number;
  seca: number;
  pitches: pitches;
  error: number;
  pitchCount: number;
  lob: number;
  inn1: number;
  bfStart: number;
  oab: number;
  cWp: number;
  bbpa: number;
  bk: number;
  obp: number;
  kbb: number;
  outcome: outcome;
  ip2: number;
  inPlay: inPlay;
  xbh: number;
  ip1: number;
  games: games;
  inn2: number;
  tc: number;
  onbase: onbase;
  tp: number;
  avg: string;
  popup: number;
  ops: number;
  abhr: number;
  linedrive: number;
  assists: assists;
  teamLob: number;
  whip: number;
  pb: number;
  abRisp: number;
  hitRisp: number;
  era: number;
  gbfb: number;
  slg: number;
  oba: number;
  k9: number;
  steal: steal;
  errors: errors;
  bf: number;
  babip: number;
  runs: runs;
  iso: number;
  po: number;
  wp: number;
  bfIp: number;
  lobRisp2out: number;
}

interface pitches {
  ktotal: number;
  perIp: number;
  perBf: number;
  perStart: number;
  count: number;
  btotal: number;
}

interface pitching {
  win: win;
  loss: loss;
  save: save;
  hold: object[];
  overall: overall;
  starters: starters;
  bullpen: bullpen;
}

interface players {
  id: string;
  position: string;
  primaryPosition: string;
  jerseyNumber: string;
  fullName: string;
  status: string;
  lastName: string;
  suffix: string;
  statistics: statistics;
  firstName: string;
  preferredName: string;
}

interface positions {
  error: number;
  pb: number;
  rf: number;
  fpct: number;
  a: number;
  cWp: number;
  tp: number;
  errors: errors;
  games: games;
  steal: steal;
  dp: number;
  position: string;
  po: number;
  inn1: number;
  assists: assists;
  tc: number;
  inn2: number;
}

interface probablePitcher {
  firstName: string;
  jerseyNumber: string;
  id: string;
  era: number;
  lastName: string;
  fullName: string;
  win: number;
  loss: number;
  preferredName: string;
}

interface reviews {
  home: home;
  away: away;
}

interface roster {
  fullName: string;
  position: string;
  preferredName: string;
  id: string;
  suffix: string;
  lastName: string;
  jerseyNumber: string;
  status: string;
  primaryPosition: string;
  firstName: string;
}

interface runs {
  total: number;
  unearned: number;
  earned: number;
  ir: number;
  ira: number;
  bqr: number;
  bqra: number;
}

interface save {
  save: number;
  blownSave: number;
  preferredName: string;
  id: string;
  loss: number;
  jerseyNumber: string;
  status: string;
  primaryPosition: string;
  win: number;
  hold: number;
  firstName: string;
  position: string;
  fullName: string;
  lastName: string;
}

interface scoring {
  type: string;
  number: number;
  sequence: number;
  runs: number;
  hits: number;
  errors: number;
}

interface starters {
  lob: number;
  bfIp: number;
  wp: number;
  games: games;
  whip: number;
  ip1: number;
  gofo: number;
  babip: number;
  kbb: number;
  inPlay: inPlay;
  onbase: onbase;
  bf: number;
  obp: number;
  outs: outs;
  pitchCount: number;
  steal: steal;
  bk: number;
  k9: number;
  timesThroughOrder: number;
  slg: number;
  ip2: number;
  oba: number;
  outcome: outcome;
  runs: runs;
  pitches: pitches;
  gbfb: number;
  oab: number;
  bfStart: number;
  era: number;
}

interface startingPitcher {
  lastName: string;
  firstName: string;
  jerseyNumber: string;
  id: string;
  loss: number;
  preferredName: string;
  fullName: string;
  win: number;
  era: number;
}

interface statistics {
  pitching: pitching;
  fielding: fielding;
  hitting: hitting;
}

interface steal {
  caught: number;
  stolen: number;
  pickoff: number;
  pct: number;
}

interface timeZones {
  venue: string;
  home: string;
  away: string;
}

interface venue {
  state: string;
  id: string;
  surface: string;
  country: string;
  location: location;
  address: string;
  fieldOrientation: string;
  stadiumType: string;
  name: string;
  city: string;
  zip: string;
  timeZone: string;
  market: string;
  capacity: number;
}

interface weather {
  currentConditions: currentConditions;
  forecast: forecast;
}

interface win {
  loss: number;
  hold: number;
  preferredName: string;
  jerseyNumber: string;
  save: number;
  fullName: string;
  firstName: string;
  position: string;
  primaryPosition: string;
  id: string;
  blownSave: number;
  lastName: string;
  status: string;
  win: number;
}

interface wind {
  speedMph: number;
  direction: string;
}

