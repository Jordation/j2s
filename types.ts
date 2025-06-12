interface agility {
  yaw: number;
  acceleration: acceleration;
  vmaxSteerMax: number;
  deceleration: deceleration;
  v0SteerMax: number;
  kvSteerMax: number;
  pitch: number;
  roll: number;
}

interface deceleration {
  main: number;
}

interface rentalPercentDays {
  duration1: number;
  duration3: number;
  duration7: number;
  duration30: number;
}

interface ship {
  data: data;
  meta: meta;
}

interface quantum {
  quantumSpoolTime: number;
  quantumFuelCapacity: number;
  quantumRange: number;
  quantumSpeed: number;
}

interface data {
  fuel: fuel;
  armor: armor;
  quantum: quantum;
  uuid: string;
  skus: object[];
  personalInventory: number;
  name: string;
  vehicleInventory: number;
  version: string;
  emission: emission;
  id: number;
  loaner: any[];
  description: string;
  msrp: number;
  size: string;
  pledgeUrl: string;
  mass: number;
  cargoCapacity: number;
  agility: agility;
  insurance: insurance;
  sizeClass: number;
  manufacturer: manufacturer;
  type: string;
  updatedAt: string;
  productionNote: string;
  chassisId: number;
  health: number;
  shops: object[];
  shieldHp: number;
  sizes: sizes;
  slug: string;
  parts: object[];
  foci: string[];
  productionStatus: string;
  crew: crew;
  speed: speed;
  components: object[];
  className: string;
}

interface parts {
  displayName: string;
  damageMax: number;
  children: object[];
  name: string;
}

interface priceRange {
  min: number;
  max: number;
}

interface speed {
  scmToZero: number;
  maxToZero: number;
  reverse: number;
  max: number;
  scm: number;
  zeroToScm: number;
  zeroToMax: number;
}

interface emission {
  emMax: number;
  ir: number;
  emIdle: number;
}

interface skus {
  title: string;
  price: number;
  available: number;
  importedAt: string;
}

interface fuel {
  capacity: number;
  intakeRate: number;
  usage: usage;
}

interface armor {
  signalInfrared: number;
  signalCrossSection: number;
  damagePhysical: number;
  damageEnergy: number;
  signalElectromagnetic: number;
  damageDistortion: number;
  damageThermal: number;
  damageBiochemical: number;
  damageStun: number;
}

interface children {
  damageMax: number;
  children: object[];
  name: string;
  displayName: string;
}

interface crew {
  min: number;
  max: null;
  weapon: number;
  operation: null;
}

interface manufacturer {
  code: string;
  name: string;
}

interface loaner {
  name: string;
  link: string;
  version: string;
}

interface meta {
  processedAt: string;
  validRelations: string[];
}

interface items {
  buyable: boolean;
  uuid: string;
  basePrice: number;
  maxDiscount: number;
  maxInventory: number;
  maxPremium: number;
  optimalInventory: number;
  priceCalculated: number;
  autoRestock: boolean;
  subType: string;
  rentable: boolean;
  name: string;
  inventory: number;
  refreshRate: number;
  type: string;
  version: string;
  basePriceOffset: number;
  rentalPriceDays: rentalPriceDays;
  autoConsume: boolean;
  sellable: boolean;
  priceRange: priceRange;
  rentalPercentDays: rentalPercentDays;
}

interface shops {
  position: string;
  version: string;
  items: object[];
  name: string;
  profitMargin: number;
  link: string;
  uuid: string;
  nameRaw: string;
}

interface insurance {
  claimTime: number;
  expediteTime: number;
  expediteCost: number;
}

interface acceleration {
  mainG: number;
  retroG: number;
  maneuveringG: number;
  vtol: number;
  xAxis: null;
  zAxis: null;
  vtolG: number;
  main: number;
  retro: number;
  maneuvering: number;
  yAxis: null;
}

interface sizes {
  length: number;
  beam: number;
  height: number;
}

interface components {
  componentSize: string;
  type: string;
  mounts: number;
  category: string;
  size: string;
  details: string;
  manufacturer: string;
  name: string;
  quantity: number;
  componentClass: string;
}

interface rentalPriceDays {
  duration30: number;
  duration1: number;
  duration3: number;
  duration7: number;
}

interface usage {
  vtol: number;
  main: number;
  maneuvering: number;
  retro: number;
}

