import { TextInput } from '@mantine/core';

interface SearchBarProps {
    label: string;
    description?: string;
    placeholder?: string
}

function SearchBar(props: SearchBarProps) {
    const { label, description, placeholder } = props;
    return (
        <TextInput
            label={label}
            description={description}
            placeholder={placeholder}
        />
    );
}
export default SearchBar

