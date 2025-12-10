// types.ts

export type Data = {
  id: string;
  name: string;
  address: {
    street: string;
    city: string;
    state: string;
    zip: string;
  };
  phone: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
};

export type DataParserError = {
  code: number;
  message: string;
  details: {
    field: string;
    message: string;
  }[];
};

export type ParsedData = {
  id: string;
  name: string;
  address: {
    street: string;
    city: string;
    state: string;
    zip: string;
  };
  phone: string;
  email: string;
};

export type ParserConfig = {
  fields: {
    id?: string;
    name?: string;
    address?: {
      street?: string;
      city?: string;
      state?: string;
      zip?: string;
    };
    phone?: string;
    email?: string;
  };
  errorHandler?: (error: DataParserError) => void;
};

export type ParserCallback = (error: DataParserError | null, data: ParsedData[]) => void;