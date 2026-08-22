export interface KataSummary {
  id: string;
  slug: string;
  title: string;
  focus: string;
  signature: string;
  evaluatorStatus: string;
  level: string;
  language: string;
  tags: string[];
}

export interface CategorySummary {
  id: string;
  title: string;
  learningGoal: string;
  katas: KataSummary[];
}

export interface StageSummary {
  id: string;
  title: string;
  level: string;
  categories: CategorySummary[];
}

export interface Track {
  id: string;
  title: string;
  description: string;
  stages: StageSummary[];
  kataCount: number;
}

export interface KataContent {
  id: string;
  slug: string;
  kataGo: string;
  kataTest: string;
  buggyKata: string;
  readme: string;
  json: string;
}

export interface KataDetail {
  kata: {
    id: string;
    slug: string;
    title: string;
    focus: string;
    signature: string;
    evaluatorStatus: string;
    level: string;
    language: string;
    rules: string[];
    flashcards: any[];
    quizQuestions: any[];
  };
  content: KataContent;
}

export interface RunResult {
  status: string;
  passed: boolean;
  failedTests: string[];
  output: string;
  evaluatorError: string;
  duration: string;
}

export interface ProgressState {
  attempts: Record<string, {
    passes: number;
    lastResult?: string;
  }>;
}

export interface AppStatus {
  kataCount: number;
  stageCount: number;
  hasRunner: boolean;
  syncReady: boolean;
  contentDir: string;
}
