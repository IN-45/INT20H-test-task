const TYPES = ['image/jpeg', 'image/png'];

export const getStylesByState = (state: boolean, type: string) => {
  const styleClasses = state
    ? '[&_button]:bg-black-secondary [&_button]:hover:bg-orange-main [&_button]:hover:text-black-secondary'
    : 'border-dashed border-2 border-black-secondary text-black-main hover:text-orange-main hover:border-orange-main';
  const typeClasses = type === 'avatar' ? '[&_p]:opacity-70 [&_p]:hover:opacity-100' : '';
  return `${styleClasses} ${typeClasses}`;
};

export const getStylesByType = (type: string) => {
  switch (type) {
    case 'avatar':
      return 'w-[86px] h-[86px] rounded-full';
    case 'document':
      return 'w-full min-h-[188px] h-fit rounded-lg gap-y-4';
  }
};

export const getStylesByError = (isError: boolean) => {
  return isError ? 'border-dashed border-4 border-red-secondary box-content' : 'cursor-pointer';
};

export const getTextStyles = (type: string) => {
  switch (type) {
    case 'avatar':
      return 'text-2xs font-medium w-10';
    case 'document':
      return 'font-semibold text-md p-y-1';
  }
};

export const getButtonStyles = (type: string) => {
  switch (type) {
    case 'avatar':
      return 'p-[5px]';
    case 'document':
      return 'p-2';
  }
};

export const getImgStyles = (type: string) => {
  switch (type) {
    case 'avatar':
      return 'object-fill';
    case 'document':
      return 'object-contain';
  }
};

export const checkFile = (file: File, type: string) => {
  if (file.size > 5 * 1000000) {
    return false;
  }
  if (!TYPES.includes(file.type)) {
    return type === 'document' && file.type === 'application/pdf';
  }
  return true;
};
